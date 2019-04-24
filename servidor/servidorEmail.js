const nodemailer = require("nodemailer")

export default class EmailServer(){


let trasporter = {};

dameTemplateHTML(nombreTemplate) {
    var html = new Promise((resolve, reject) => {
        let rutaTemplate = path.join('.','emails',nombreTemplate,'html.hbs')
    fs.readFile(rutaTemplate, 'utf-8', (err, datos) => {
      err ? reject(err) : resolve(datos)
    })
  })

    var subject = new Promise((resolve, reject) => {
    let rutaTemplate = path.join('.','emails',nombreTemplate,'subject.hbs')
    fs.readFile(rutaTemplate, 'utf-8', (err, datos) => {
      err ? reject(err) : resolve(datos)
    })
  })

  return Promise.all([html, subject])
}


guardaPlantilla(html, ruta, nombreArchivo) {
  if (html && ruta) {
    mkdirp(ruta, function (err) {
      if (err) {
        console.log("Error en la ruta esta")
      } else {
          let path = path.join(ruta,nombreArchivo)
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

/*
 * Fijar el medio por el cual se enviaran los correos
 * confSMTP: Es la configuracion SMTP del servidor que enviara el correo
 */
fijaTransporte(confSMTP){
    this.transporter = nodemailer.createTransport(confSMTP);;
}

enviaCorreo(){

}

}

