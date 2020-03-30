import { html } from "https://unpkg.com/lit-html?module";

import formulario from "./formulario.js";
import { MyForm } from "./form/form.js";

const dashboard = () => {
  //Elementos de inputs del fomulario
  const inputsClient = [
    {
      name: "name",
      type: "text",
      label: "Nombre",
    },
  ];

  // Como procesar los datos del submit
  const sub = (data, form) => {
    fetch("/client", {
      method: "POST",
      body: JSON.stringify(Object.fromEntries(data)),
    }).then((r) => {
      if (r.ok) {
        form.reset();
      } else {
        console.log("Error");
      }
    });
  };
  //Crear un formulario
  let formCliente = new MyForm("clients", inputsClient, sub);

  return html`
    <toast-noti titulo="Titulo" message="Prueba mensaje" duration="50000" />
    ${formCliente}
  `;
};
export { dashboard };
