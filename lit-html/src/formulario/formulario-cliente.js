import {LitElement, html} from 'lit-element';
import '@vaadin/vaadin-text-field/vaadin-text-field.js';
import '@vaadin/vaadin-form-layout/vaadin-form-layout.js';
import '@vaadin/vaadin-button/vaadin-button.js';

export class FormularioCliente extends LitElement {
  static get properties() {
    return {
      _id: {
        type: Number,
      },
      cliente: {
        type: String,
      },
      numeroFactura: {
        type: String,
      },
      ordenDeCompra: {
        type: String,
      },
    };
  }

  constructor() {
    super();
    //Nombre del cliente
    this.cliente = '';
    //Numero de factura
    this.numeroFactura = '';
    //Numero de orden de compra
    this.ordenDeCompra = '';
  }

  fijarCliente({target: {value}}) {
    this.cliente = value;
  }

  fijarNumeroFactura({target: {value}}) {
    this.numeroFactura = value;
  }

  fijarOrdenDeCompra({target: {value}}) {
    this.ordenDeCompra = value;
  }
  //TODO Probar el generador de id's
  guardarFormulario() {
    const orden = {
      cliente: this.cliente,
      numeroFactura: this.numeroFactura,
      ordenDeCompra: this.ordenDeCompra,
    };

    const event = new CustomEvent('guardar', {
      detail: {
        orden,
        error: false,
      },
    });

    this.dispatchEvent(event);
  }

  render() {
    return html`
      <vaadin-form-layout>
        <vaadin-text-field
          label="cliente"
          @change=${this.fijarCliente}
        ></vaadin-text-field>

        <vaadin-text-field
          label="Factura"
          @change=${this.fijarNumeroFactura}
        ></vaadin-text-field>

        <vaadin-text-field
          label="Orden de compra"
          @change=${this.fijarOrdenDeCompra}
        ></vaadin-text-field>

        <vaadin-button @click=${this.guardarFormulario}>Guardar</vaadin-button>
      </vaadin-form-layout>
    `;
  }
}

customElements.define('formulario-cliente', FormularioCliente);
