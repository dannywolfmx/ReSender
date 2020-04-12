import { html, render } from "https://unpkg.com/lit-html?module";
import "./arte-icons.js";
const style = html`
  <style>
    :host {
      display: flex;
      padding: 8px 0px;
      align-content: center;
      flex-direction: row;
      justify-content: center;
    }

    :host(:hover) {
      background: #6777ef;
      color: white;
    }
    p {
      margin: 0px;
      display: flex;
      align-items: center;
      flex: 1;
      justify-content: center;
    }
  </style>
`;

export class MenuItem extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
  }
  connectedCallback() {
    this.update();
  }

  template() {
    return html`
      ${style}
      <span is="arte-icons" icon="contactless"> </span>
      <p>${this.title}</p>
    `;
  }

  update() {
    render(this.template(), this.root, { eventContext: this });
  }
}
customElements.define("menu-item", MenuItem);
