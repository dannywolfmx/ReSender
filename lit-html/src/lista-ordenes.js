import {LitElement, html} from 'lit-element';
import './tarjeta/info-cliente.js';

export class ListaOrdenes extends LitElement {
  static get properties() {
    return {
      listaOrdenes: {
        type: Array,
      },
    };
  }

  constructor() {
    super();

    this.listaOrdenes = [
      {
        cliente: {
          nombre: 'Nombre de prueba',
        },
        orden: 1234,
        factura: 'FAC234',
      },
    ];
  }

  render() {
    return html`
      ${this.listaOrdenes.map(
        orden =>
          html`
            <info-cliente .orden=${orden}> </info-cliente>
          `,
      )}
    `;
  }
}

customElements.define('lista-ordenes', ListaOrdenes);
