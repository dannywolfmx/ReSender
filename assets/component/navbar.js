import { html, render } from "https://unpkg.com/lit-html?module";

const style = html`
  <style>
    #navbar {
      background: #6777ef;
      display: flex;
      padding: 8px;
      height: 56px;
      flex-direction: row;
    }
    #title {
      color: white;
    }
    nav ::slotted(input) {
      display: flex;
      color: red;
      margin-left: 24px !important;
    }
  </style>
`;

export class NavBar extends HTMLElement {
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
      <nav id="navbar">
        <slot name="icon"></slot>
        <slot name="title" id="title"></slot>
        <slot name="menu" id="menu"></slot>
        <slot name="search" id="search"></slot>
      </nav>
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}

customElements.define("nav-bar", NavBar);
