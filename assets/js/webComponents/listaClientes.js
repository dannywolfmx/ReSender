class ListaClientes extends HTMLElement{
	constructor(){
		super()
	}
	
	connectedCallback(){
		this.actualiza()
	}
	
	//Imprime la lista de clientes
	_render(clientes){
		 console.log(clientes)
		 this.innerHTML= `
			${clientes.map((cliente)=> `<button>${cliente.nombre}</button>`).join('')}
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
