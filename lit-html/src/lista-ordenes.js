import {LitElement, html} from 'lit-element';
import './tarjeta/info-cliente.js';

export class ListaOrdenes extends LitElement {
  static get properties() {
    return {
      ordenes: {
        type: Array,
      },
    };
  }

  constructor() {
    super();
    this.ordenes = [];
  }

  editarOrden({detail: {orden}}) {
    console.log(orden);
  }

  eliminarOrden({detail}) {
    const evento = new CustomEvent('eliminar', {detail});
    this.dispatchEvent(evento);
  }
  //Tarjeta con informacion de la orden

  tarjetaInformacion(orden) {
    return html`
      <info-cliente
        .orden=${orden}
        @editar=${this.editarOrden}
        @eliminar="${this.eliminarOrden}"
      ></info-cliente>
    `;
  }

  render() {
    return html`
      ${this.ordenes.map(this.tarjetaInformacion.bind(this))}
    `;
  }
}

customElements.define('lista-ordenes', ListaOrdenes);
