<script>
export default {
    data() {
        return {
            figureName: "",
        };
    },

    methods: {
        async fetchFigures() {
            const resp = await fetch(`http://localhost:3000/fetchFigures?name=${this.figureName}`)
            console.log("clicking the fetch button");
            const ret = await resp.json();


            ret.forEach(entry => {
                const figure = document.createElement("li");
                figure.appendChild(document.createTextNode(entry.name));
                figure.appendChild(document.createTextNode("Birth: "));
                figure.appendChild(document.createTextNode(entry.birthdate));
                figure.appendChild(document.createTextNode("Death: "));
                figure.appendChild(document.createTextNode(entry.deathdate));
                document.getElementById("figures").appendChild(figure);

            });
        }
    }
}
</script>

<template>
    <div>
        <h2>Historical Figures</h2>
        <p>Welcome Historical Figures Timeline</p>

        <input type="text" id="figure-name" v-model="figureName">
        <button @click="fetchFigures">
            test figures
        </button>
        <ul id="figures"></ul>
    </div>
</template>