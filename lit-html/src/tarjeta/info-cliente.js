import {LitElement, html} from 'lit-element';
import '@vaadin/vaadin-button/vaadin-button.js';

export class InfoCliente extends LitElement {
  constructor() {
    super();
  }

  static get properties() {
    /*
     * cliente type String
     * numeroFactura type String
     * ordenDeCompra type String
     */
    return {
      orden: Object,
    };
  }
  //Editar el formulario
  editar() {}
  //No eliminar, solo marcar como inactivo
  eliminar() {}

  render() {
    return html`
      <div>
        <div>Id: ${this.orden._id}</div>
        <div>Cliente: ${this.orden.cliente}</div>
        <div>Factura: ${this.orden.numeroFactura}</div>
        <div>Orden: ${this.orden.ordenDeCompra}</div>
      </div>
      <div>
        <vaadin-button @click="${this.editar}">
          Editar
        </vaadin-button>
        <vaadin-button @click="${this.eliminar}">
          Eliminar
        </vaadin-button>
      </div>
    `;
  }
}

customElements.define('info-cliente', InfoCliente);
