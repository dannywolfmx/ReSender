// Modules to control application life and create native browser window
const {
  app,
  BrowserWindow,
  ipcMain,
  shell
} = require('electron');
//Trabajar con rutas entre plataformas
const path = require('path');

const fs = require('fs');
const directorioOC = path.join('.',"OrdenesDeCompra")
const juice = require('juice');
const dot = require('dot')
const mkdirp = require('mkdirp');

const ClienteEmail = require('./servidor/servidorEmail.js')

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow

let clienteEmail = new ClienteEmail();

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
    
    //Crear el directorio para ordenes de compra si no existe 
  if (!fs.existsSync(directorioOC)) {
    fs.mkdirSync(directorioOC);
  }

  ipcMain.on("copiarArchivo", function (event, archivo) {
    if (archivo) {
        let directorio = path.join(directorioOC, archivo.oc);
      if (!fs.existsSync(directorio)) {
        fs.mkdirSync(directorio);
      }
        console.log(archivo.ruta)
        let rutaArchivo = path.join(directorio,archivo.nombre);
      fs.createReadStream(archivo.ruta).pipe(fs.createWriteStream(rutaArchivo));
    }
  })

    ipcMain.on("abreArchivo",function(event,ruta){
    let rutaArchivo = path.join(app.getAppPath(),ruta)
    shell.showItemInFolder(rutaArchivo)
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



function getDirectories(ruta) {
    return fs.readdirSync(ruta).filter(function (file) {
        let rutaArchivo = path.join(ruta,file)
    return fs.statSync(rutaArchivo).isDirectory();
  });
}



ipcMain.on("enviarCorreo", function (event, templateCorreo) {
    let clienteEmail = new ClienteEmail();
    clienteEmail.enviarCorreo(templateCorreo);
})
