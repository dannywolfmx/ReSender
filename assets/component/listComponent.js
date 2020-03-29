import { html } from "https://unpkg.com/lit-html?module";

const list = (list) => {
  return html`
    <ul class="list-group">
      ${list.map((text) => html` <li class="list-group-item">${text}</li> `)}
    </ul>
  `;
};

export { list };
