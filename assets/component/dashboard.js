import { html, render } from "https://unpkg.com/lit-html?module";
import "./navbar.js";
import "./marcadores.js";
import "./lateral-menu.js";
import "./bootstrap-input.js";

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
    #lateral {
    }
    .lateralFull {
      display: block;
      flex: 1 250px;
    }

    .lateralMin {
      display: none;
      flex: 0;
    }
    #content {
      display: flex;
      flex-direction: column;
    }

    .contentMin {
      width: calc(100% - 250px);
    }

    .contentFull {
      width: 100%;
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

  ocultarMenu(e) {
    let content = this.root.getElementById("content");
    let lateral = this.root.getElementById("lateral");

    //Ocultar o mostrar menu
    content.classList.toggle("contentMin");
    lateral.classList.toggle("lateralMin");

    content.classList.toggle("contentFull");
    lateral.classList.toggle("lateralFull");
  }

  _template() {
    return html`
      ${style}
      <lateral-menu id="lateral" class="lateralFull"></lateral-menu>
      <div id="content" class="contentMin">
        <nav-bar>
          <button @click=${this.ocultarMenu} slot="icon">Ocultar</button>
          <h1 slot="title">Titulo</h1>
          <bootstrap-input label="Prueba" slot="search"></bootstrap-input>
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
