import {LitElement, html} from 'lit-element';
import './tarjeta/info-cliente.js';

export class ListaOrdenes extends LitElement {
  static get properties() {
    return {
      lista: {
        type: Array,
      },
    };
  }

  constructor() {
    super();

    this.listaOrdenes = [];
  }
  updated(e) {
    console.log(e);
  }
  render() {
    return html`
      ${this.lista.map(orden => {
        return html`
          ${console.log(orden)}
          <info-cliente .orden=${orden}></info-cliente>
        `;
      })}
    `;
  }
}

customElements.define('lista-ordenes', ListaOrdenes);
