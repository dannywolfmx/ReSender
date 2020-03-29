import { html, render } from "https://unpkg.com/lit-html?module";

import { card } from "./cardComponent.js";
import { list } from "./listComponent.js";

class Clients extends HTMLElement {
  list = [];
  constructor() {
    super();
    this.root = this;
    this._dameCliente();
    this._update();
  }
  _dameCliente() {
    fetch("/clients")
      .then((r) => {
        return r.json();
      })
      .then((json) => {
        json.map((e) => {
          this.list.push(e.name);
        });
        this._update();
      });
  }
  _template() {
    return html` ${card("Clientes", list(this.list))} `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}

customElements.define("list-clients", Clients);
