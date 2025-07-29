

// Custom-Element my-element anlegen
export default class OtsTitle extends HTMLElement {
	
    constructor() {
        super();
        
        this.count = 0;
        this.btn = document.createElement('button');
        this.btn.textContent = 'Count up!';

        this.announce = document.createElement('div');
        this.announce.setAttribute('role', 'status');
        this.announce.textContent = `Clicked ${this.count} times!`

        this.append(this.btn);        
        this.append(this.announce);

        this.btn.addEventListener('click', this);
    }

    handleEvent(event) {
        this.count++;
        this.announce.textContent = `Clicked ${this.count} times!`
    }

    connectedCallback() {
    }

}

customElements.define('ots-title', OtsTitle);
