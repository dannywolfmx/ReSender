class DropArea extends HTMLElement{
	constructor(){
		super()
		this.attachShadow({mode:'open'})
	}
	connectedCallback(){
		this._render()		
		
		//Evento que indica que el elemento fue arrastrado al formulario
		this.addEventListener('drop',(e)=>{
			e.preventDefault()
			let items = e.dataTransfer.items
			if(items){
				let archivos = this._recuperaArchivos(items)
				console.log(archivos)
			}else{
				console.log("No se encontraron archivos")
			}
			console.log('drop')
		})

		//Evento para mostrar que el elemento esta siendo arrastrado
		//al formulario
		this.addEventListener('dragover',(e)=>{
			e.preventDefault()
			//Prograr animacion de drag y deteccion de archivos validos
		})
	}

	_render(){
		let style = `
			<style>
				#dropArea{
					height:100px;
					width: 100%;
					margin:10px;
					border:1px solid black;
				}
				.file{
					height:50px;
					width:50px;
					border:1px solid black;
					display:inline-block;
				}
			</style>
		`
		this.shadowRoot.innerHTML = `
			${style}
			<div id="dropArea">
			</div>
		`
	}
	
	//Recuperar archivos de un dataTransfer
	_recuperaArchivos(items){
		let archivos = []
		
		for(let i=0; i< items.length; i++){
			//Detectar si de tipo file
			if(items[i].kind === 'file'){
				let archivo = items[i].getAsFile();
				this._crearArchivo(archivo.name)
			}
		}
		return archivos
	}
	
	_crearArchivo(nombre){
		let contenedor = document.createElement("div")
		contenedor.className = "file"
		contenedor.innerText = nombre
		let dropArea = this.shadowRoot.getElementById('dropArea')
		dropArea.appendChild(contenedor)
	}
}

customElements.define("drop-area", DropArea )
