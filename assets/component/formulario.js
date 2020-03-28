import { html } from "https://unpkg.com/lit-html?module";

const app = () => {
  return html`
    <div class="card">
      <div class="card-body">
        <form>
          <div class="form-group">
            <label for="exampleInputEmail1">Nombre</label>
            <input
              type="text"
              class="form-control"
              id="exampleInputEmail1"
              aria-describedby="emailHelp"
            />
          </div>
          <button type="submit" class="btn btn-primary">Crear</button>
        </form>
      </div>
    </div>
  `;
};

export { app };
