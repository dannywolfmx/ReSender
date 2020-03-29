import { html } from "https://unpkg.com/lit-html?module";

const card = (titulo, elemento) => {
  return html`
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">${titulo}</h5>
        ${elemento}
      </div>
    </div>
  `;
};

export { card };
