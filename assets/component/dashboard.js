import { html, render } from "https://unpkg.com/lit-html?module";

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
export class MyDashboard extends HTMLElement {
  constructor() {
    super();
    this.inputsClient = JSON.stringify(inputsClient);
    this.inputsOrder = JSON.stringify(inputsOrder);
  }

  connectedCallback() {
    this._update();
    console.log(new POSTForm());
  }
  _template() {
    return html`
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">"Crear orders"</h5>
          <post-form
            nameform="orders"
            url="/order"
            inputs=${this.inputsOrder}
          ></post-form>
        </div>
      </div>
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">"Crear client"</h5>
          <post-form
            nameform="clients"
            url="/client"
            inputs=${this.inputsClient}
          ></post-form>
        </div>
      </div>
    `;
  }
  _update() {
    render(this._template(), this, { eventContext: this });
  }
}
customElements.define("my-dashboard", MyDashboard);
