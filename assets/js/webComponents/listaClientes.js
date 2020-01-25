class ListaClientes extends HTMLElement{
	constructor(){
		super()
		this.clientes = ["khdlskajd","akodjlkasjd"]
	}
	
	connectedCallback(){
		this.actualiza()
	}
	
	//Imprime la lista de clientes
	_render(){
		 this.innerHTML= `
		 	<ul class="list-group list-group-horizontal">
				${
					this.clientes.map(cliente =>(
						`<li class="list-group-item">${cliente.nombre}</li>`
					)).join('')
				}
			</ul>
		`
	}

	actualiza(){
		fetch('clientes').then((r)=>{
			return r.json()
		}).then(c => {
			this.clientes = c
			//Mandrar a renderizar el contenido
			this._render()
		}).catch((e)=>{
			console.log("Error al obtener clientes: ",e)
		})
	}
	
}

customElements.define("lista-clientes",ListaClientes)
