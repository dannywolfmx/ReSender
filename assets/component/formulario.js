import { html } from "https://unpkg.com/lit-html?module";
import { card } from "./cardComponent.js";

const guardar = (e) => {
  e.preventDefault();
  let formulario = document.getElementById("crearCliente");
  let data = new FormData(formulario);
  fetch("/client", {
    method: "POST",
    body: JSON.stringify(Object.fromEntries(data)),
  }).then((r) => {
    if (r.ok) {
      formulario.reset();
    } else {
      console.log("Error");
    }
  });
};

const form = html`
  <form id="crearCliente">
    <div class="form-group">
      <label for="exampleInputEmail1">Nombre</label>
      <input
        type="text"
        class="form-control"
        id="exampleInputEmail1"
        name="name"
        aria-describedby="emailHelp"
      />
    </div>
    <button type="submit" class="btn btn-primary" @click=${guardar}>
      Crear
    </button>
  </form>
`;

const tarjeta = () => {
  return card("Crear cliente", form);
};

export default tarjeta;
