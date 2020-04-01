import { html } from "https://unpkg.com/lit-html?module";
const dashboard = () => {
  //Elementos de inputs del fomulario
  const inputsClient = [
    {
      name: "name",
      type: "text",
      label: "Nombre",
    },
  ];

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
      hide: true,
      value: 7,
    },
  ];
  let serialClientInput = JSON.stringify(inputsClient);
  return html`
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">"Crear orders"</h5>
        <post-form
          nameform="orders"
          url="/order"
          inputs=${JSON.stringify(inputsOrder)}
        ></post-form>
      </div>
    </div>
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">"Crear client"</h5>
        <post-form
          nameform="clients"
          url="/client"
          inputs=${serialClientInput}
        ></post-form>
      </div>
    </div>
  `;
};
export { dashboard };
