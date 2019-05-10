import {LitElement, html} from 'lit-element';

export class InfoCliente extends LitElement {
  constructor() {
    super();
  }

  static get properties() {
    return {
      orden: Object,
    };
  }

  render() {
    return html`
      <div>${this.orden.cliente.nombre}</div>
    `;
  }
}

customElements.define('info-cliente', InfoCliente);
