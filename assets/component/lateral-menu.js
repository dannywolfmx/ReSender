import { html, render } from "https://unpkg.com/lit-html?module";
import "./menu-item.js";
const style = html`
  <style>
    :host {
      flex: 1 250px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.03);
    }

    #menu {
    }
    #title > h1 {
      height: 72px;
      background: #6777ef;
      margin: 0px;
    }
  </style>
`;

export class LateralMenu extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
  }
  connectedCallback() {
    this._update();
  }
  _template() {
    return html`
      ${style}
      <div id="title">
        <h1>Titulo</h1>
      </div>
      <div id="menu">
        <menu-item>Prueba</menu-item>
        <menu-item>Prueba</menu-item>
        <menu-item>Prueba</menu-item>
        <menu-item>Prueba</menu-item>
      </div>
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}

customElements.define("lateral-menu", LateralMenu);
