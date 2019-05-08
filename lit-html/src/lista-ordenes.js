import {LitElement, html} from 'lit-element';

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
        cliente: 'Prueba',
        orden: 1234,
        factura: 'FAC234',
      },
    ];
  }

  render() {
    return html`
      <ul>
        ${this.listaOrdenes.map(
          i =>
            html`
              <li>${i.factura}</li>
            `,
        )}
      </ul>
    `;
  }
}

customElements.define('lista-ordenes', ListaOrdenes);
