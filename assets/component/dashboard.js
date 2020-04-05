import { html, render } from "https://unpkg.com/lit-html?module";
import "./navbar.js";
import "./marcadores.js";
import "./lateral-menu.js";

const style = html`
  <style>
    :host {
      margin: 0;
      min-height: 100%;
      display: flex;
      flex-flow: row wrap;
      text-align: center;
      flex: 1;
    }
    #content {
      display: flex;
      flex-direction: column;
      width: calc(100%-250px);
    }
  </style>
`;
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
    this.root = this.attachShadow({ mode: "open" });
  }

  connectedCallback() {
    this._update();
  }
  _template() {
    return html`
      ${style}
      <lateral-menu id="lateral"></lateral-menu>
      <div id="content">
        <nav-bar>
          <h1 slot="title">Titulo</h1>
          <input placeholder="search" slot="search" />
        </nav-bar>

        <my-marcadores>
          <div class="card" slot="marcadores">
            <div class="card-body">
              <h5 class="card-title">"Crear orders"</h5>
              <post-form
                nameform="orders"
                url="/order"
                inputs=${this.inputsOrder}
              ></post-form>
            </div>
          </div>
          <div class="card" slot="marcadores">
            <div class="card-body">
              <h5 class="card-title">"Crear client"</h5>
              <post-form
                nameform="clients"
                url="/client"
                inputs=${this.inputsClient}
              ></post-form>
            </div>
          </div>
        </my-marcadores>
      </div>
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}
customElements.define("my-dashboard", MyDashboard);
