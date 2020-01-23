class ListaClientes extends HTMLElement{
	constructor(){
		super()
		this.attachShadow({mode:'open'})
	}
	
	connectedCallback(){
		this.actualiza()
	}
	
	//Imprime la lista de clientes
	_render(clientes){
		 this.shadowRoot.innerHTML= `
			<ul>
				${clientes.map((cliente)=> `<li>${cliente.nombre}</li>`).join('')}
			</ul>
		`
	}
	
	actualiza(){
		fetch('clientes').then((r)=>{
			return r.json()
		}).then(this._render.bind(this)).catch((e)=>{
			console.log("Error al obtener clientes: ",e)
		})
	}
	
}

customElements.define("lista-clientes",ListaClientes)
