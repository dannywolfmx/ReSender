import { html } from "https://unpkg.com/lit-html?module";

import formulario from "./formulario.js";
import { MyForm } from "./form/form.js";

class Formulario extends MyForm {
  constructor(nameForm, inputs, url) {
    super();
    // Como procesar los datos del submit
  }
}

const dashboard = () => {
  //Elementos de inputs del fomulario
  const inputsClient = [
    {
      name: "name",
      type: "text",
      label: "Nombre",
    },
  ];

  //Crear un formulario
  let formCliente = new MyForm("clients", inputsClient, "/client");

  const inputsOrder = [
    {
      name: "number",
      type: "text",
      label: "numero",
    },
    {
      name: "invoice",
      type: "text",
      label: "Factura",
    },
    {
      name: "ClientID",
      type: "number",
      label: "Client ID",
    },
  ];
  let formOrder = new MyForm("order", inputsOrder, "/order");

  return html`
    <toast-noti titulo="Titulo" message="Prueba mensaje" duration="50000" />
    ${formCliente} ${formOrder}
  `;
};
export { dashboard };
