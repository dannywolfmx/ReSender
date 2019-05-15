import {LitElement, html} from 'lit-element';
import './tarjeta/info-cliente.js';
import {repeat} from 'lit-html/directives/repeat';

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
    const evento = new CustomEvent('editar', {detail: {_id: orden._id}});
    this.dispatchEvent(evento);
  }

  eliminarOrden({detail}) {
    const evento = new CustomEvent('eliminar', {detail});
    this.dispatchEvent(evento);
  }
  //Tarjeta con informacion de la orden
  tarjetaInformacion(orden) {
    //La orden de compra se tiene que enviar un "clon" para que lit detecte el cambio
    return html`
      <info-cliente
        .orden=${{...orden}}
        @editar=${this.editarOrden}
        @eliminar="${this.eliminarOrden}"
      ></info-cliente>
    `;
  }

  render() {
    return html`
      ${this.tarjetaInformacion(this.ordenes[0])}
      ${repeat(
        this.ordenes,
        orden => orden._id,
        this.tarjetaInformacion.bind(this),
      )}
    `;
  }
}

customElements.define('lista-ordenes', ListaOrdenes);
