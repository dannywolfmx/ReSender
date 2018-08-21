// Modules to control application life and create native browser window
const {
  app,
  BrowserWindow,
  ipcMain
} = require('electron')
const fs = require('fs');
let nodemailer = require("nodemailer")
const directorioOC = "OrdenesDeCompra"
var juice = require('juice');
var dot = require('dot')

var mkdirp = require('mkdirp');

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow

function createWindow() {
  // Create the browser window.
  mainWindow = new BrowserWindow({
    width: 800,
    height: 500,
    show: false,
    frame: true,
    fullscreenable: false,
    resizable: true,
    transparent: false,
    //Mostrar icono en la barra de tareas
    skipTaskbar: false,
    'node-integration': false,
    webPreferences: {
      nodeIntegrationInWorker: true,
      devTools: true
    }
  })

  // mainWindow.setMenu(null)

  // and load the index.html of the app.
  mainWindow.loadURL(`file://${__dirname}/polymer/index.html`)

  // Open the DevTools.
  // mainWindow.webContents.openDevTools()

  // Emitted when the window is closed.
  mainWindow.on('closed', function () {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    mainWindow = null
  })

  mainWindow.once('ready-to-show', () => {
    mainWindow.show(); // Thumbbar is not showing

  });

  if (!fs.existsSync(directorioOC)) {
    fs.mkdirSync(directorioOC);
  }

  ipcMain.on("copiarArchivo", function (event, archivo) {
    if (archivo) {
      let directorio = directorioOC + "\\" + archivo.oc
      if (!fs.existsSync(directorio)) {
        fs.mkdirSync(directorio);
      }
      fs.createReadStream(archivo.ruta).pipe(fs.createWriteStream(`${directorio}\\${archivo.nombre}`));
    }
  })

}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow)

// Quit when all windows are closed.
app.on('window-all-closed', function () {
  // On OS X it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', function () {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (mainWindow === null) {
    createWindow()
  }
})

function dameTemplateHTML(nombreTemplate) {
  var html = new Promise((resolve, reject) => {
    fs.readFile(`./emails/${nombreTemplate}/html.hbs`, 'utf-8', (err, datos) => {
      err ? reject(err) : resolve(datos)
    })
  })

  var subject = new Promise((resolve, reject) => {
    fs.readFile(`./emails/${nombreTemplate}/subject.hbs`, 'utf-8', (err, datos) => {
      err ? reject(err) : resolve(datos)
    })
  })

  return Promise.all([html, subject])
}

function guardaPlantilla(html, ruta, nombreArchivo) {
  if (html && ruta) {
    mkdirp(ruta, function (err) {
      if (err) {
        console.log("Error en la ruta esta")
      } else {
        let path = `${ruta}/${nombreArchivo}.inline`
        if(!fs.existsSync(path)){
          fs.writeFile(path, html, function(err){
            if(err){
              return -1
            }else{
              console.log("Archivo guardado")
            }
          })
        }
      }
    })
  }
}

function getDirectories(path) {
  return fs.readdirSync(path).filter(function (file) {
    return fs.statSync(path+'/'+file).isDirectory();
  });
}



ipcMain.on("enviarCorreo", function (event, templateCorreo) {
  let transporter = nodemailer.createTransport(templateCorreo.smtp)
  let localsTemplate = templateCorreo.send.locals
  let mailOptions = templateCorreo.send.message


/*   dameTemplateHTML(templateCorreo.send.template).then(data =>{
    guardaPlantilla(juice(data[0]),"./templateEmail",templateCorreo.send.template)
  }) */
  dameTemplateHTML(templateCorreo.send.template)
    .then(data => {
      let html = dot.template(data[0])
      html = html(localsTemplate)

      let subject = dot.template(data[1])
      subject = subject(localsTemplate)
      let inline = juice(html)

      return Promise.all([inline, subject])
    }).then(
      datos => {
        mailOptions.html = datos[0]
        mailOptions.subject = datos[1]

        transporter.sendMail(mailOptions, (error, info) => {
          if (error) {
            console.log(error)
            event.sender.send("respuestaEnvioCorreo", false)
            return -1
          } else {
            event.sender.send("respuestaEnvioCorreo", true)
          }
        })
      }
    )
})