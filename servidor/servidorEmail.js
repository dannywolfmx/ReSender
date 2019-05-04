'use strict';

const nodemailer = require('nodemailer');
const juice = require('juice');
const dot = require('dot');
const path = require('path');
const fs = require('fs');

class ClienteCorreo {
  dameTemplateHTML(nombreTemplate) {
    var html = new Promise((resolve, reject) => {
      let rutaTemplate = path.join('.', 'emails', nombreTemplate, 'html.hbs');
      fs.readFile(rutaTemplate, 'utf-8', (err, datos) => {
        err ? reject(err) : resolve(datos);
      });
    });

    var subject = new Promise((resolve, reject) => {
      let rutaTemplate = path.join(
        '.',
        'emails',
        nombreTemplate,
        'subject.hbs',
      );
      fs.readFile(rutaTemplate, 'utf-8', (err, datos) => {
        err ? reject(err) : resolve(datos);
      });
    });

    return Promise.all([html, subject]);
  }

  guardaPlantilla(html, ruta, nombreArchivo) {
    if (html && ruta) {
      mkdirp(ruta, function(err) {
        if (err) {
          console.log('Error en la ruta esta');
        } else {
          let path = path.join(ruta, nombreArchivo);
          if (!fs.existsSync(path)) {
            fs.writeFile(path, html, function(err) {
              if (err) {
                return -1;
              } else {
                console.log('Archivo guardado');
              }
            });
          }
        }
      });
    }
  }

  /*
   * Fijar el medio por el cual se enviaran los correos
   * confSMTP: Es la configuracion SMTP del servidor que enviara el correo
   */
  fijaTransporte(confSMTP) {
    this.transporter = nodemailer.createTransport(confSMTP);
  }

  enviarCorreo({send, smtp}) {
    this.dameTemplateHTML(send.template)
      .then(data => {
        let html = dot.template(data[0]);
        html = html(send.locals);

        let subject = dot.template(data[1]);
        subject = subject(send.locals);
        let inline = juice(html);

        return Promise.all([inline, subject]);
      })
      .then(datos => {
        send.message.html = datos[0];
        send.message.subject = datos[1];

        nodemailer
          .createTransport(smtp)
          .sendMail(send.message, (error, info) => {
            if (error) {
              console.log(error);
              event.sender.send('respuestaEnvioCorreo', false);
              return -1;
            } else {
              console.log(info);
              event.sender.send('respuestaEnvioCorreo', true);
            }
          });
      });
  }
}

module.exports = ClienteCorreo;
