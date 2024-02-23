<script>
export default {
    data() {
        return {
            figureName: "",
            figuresList: [],
            figureSum: "",
            // figureImage: HTMLDivElement,

        };
    },

    methods: {
        // Full Transparency the I figured out how to reformat the sql Date yyyy-mmmm-dddd into Mon dd, yyyy by asking ChatGPT
        async fetchSum(bioName) {
            const bioResp = await fetch(`https://newportfolio-mdwh.onrender.com/fetchBio?name=${bioName}`)
            console.log("clicking the fetch button");
            const ret = await bioResp.json();

            this.figureSum = ret;

            console.log("fethcing image for " + bioName);
            // const imageResp = await fetch(`http://localhost:10000/fetchImage?name=${bioName}`)
            const imageResp = await fetch(`https://newportfolio-mdwh.onrender.com/fetchImage?name=${bioName}`)
            console.log(imageResp);
            const imageRet = await imageResp.json();

            var tempElement = document.createElement('div');
            tempElement.innerHTML = imageRet;

            var tempTempElement = document.createElement('div');
            tempTempElement.innerHTML = tempElement.getElementsByClassName("infobox-image")[0].innerHTML;



            document.getElementById("figure-image").appendChild(tempTempElement);
            document.getElementById("figure-image").removeChild
                (
                    document.getElementById("figure-image").firstElementChild
                );
            document.getElementById("figure-image").appendChild(tempTempElement);






        },

        async fetchFigures() {
            document.getElementById("loading").hidden = false;
            // var emptyDiv = document.createElement("div");
            // document.getElementById("figure-image").appendChild(emptyDiv);

            const resp = await fetch(`https://newportfolio-mdwh.onrender.com/fetchFigures?name=${this.figureName}`)
            console.log("clicking the fetch button");
            const ret = await resp.json();

            const dateFormat = {
                year: 'numeric',
                month: 'long',
                day: 'numeric'
            };


            // document.getElementById("figures").innerHTML = "";

            ret.forEach(entry => {

                const birthdate = new Date(entry.birthdate);
                const deathdate = new Date(entry.deathdate);

                const formattedBirth = birthdate.toLocaleDateString('en-US', dateFormat);
                const formattedDeath = deathdate.toLocaleDateString('en-US', dateFormat);

                entry.birthdate = formattedBirth;
                entry.deathdate = formattedDeath;


                // const figure = document.createElement("li");
                // figure.innerHTML = `<p>${entry.name}<br>Birth-Death: ${formattedBirth} - ${formattedDeath} <button @click="fetchBio" v-model="${entry.name}">Bio</button> </p>`
                // // const bio = document.createElement("button");
                // // bio.setAttribute("v-model", entry.name);
                // // bio.addEventListener('click', fetchBio);
                // // bio.innerHTML = "Bio"
                // // figure.appendChild(bio);
                // document.getElementById("figures").appendChild(figure);

            });

            this.figuresList = ret;
            document.getElementById("loading").hidden = true;
        },


    }
}
</script>

<template id="top">
    <div class="page-wrapper">
        <h1 id="timeline-title">Timeline of Historical Figures</h1>
        <p class="timeline-description">Welcome to my timeline of historical figures!
        </p>
        <p class="timeline-description">You will be met with a list of figures who lived at the same time. Some results can
            be suprising.</p>
        <p class="timeline-description">Try "Queen Victoria", "George Washington", "Napoleon Bonaparte"</p>

        <div id="search-bar">
            <input type="text" id="search-input" v-model="figureName">
            <button id="search-button" @click="fetchFigures">
                search
            </button>
        </div>


        <p id="loading" hidden>
            Loading Figures...
        </p>

        <ul v-if="figuresList.length" id="figures">
            <p v-if="figureSum" id="figureSum">
                <!-- <button @click="fetchImage(figure.name)">See Image</button> -->
            <div id="figure-image">
                <!--<img v-bind:src="figureImage" alt="figureImage" id="figure-image"> -->
            </div>
            {{ figureSum }}
            </p>


            <li v-for="figure in figuresList" :key="figure.id">
                <div class="figure">
                    <p class="figure-name">
                        {{ figure.name }}
                    </p>

                    <p>
                        Born: {{ figure.birthdate }}
                    </p>
                    <p>
                        Died: {{ figure.deathdate }}
                    </p>
                    <a href="#top" @click="fetchSum(figure.name)">Bio</a>

                </div>



            </li>
        </ul>

    </div>
</template>


<style scoped>
template {

    scroll-behavior: smooth;
}

#loading {
    align-self: center;
}

.page-wrapper {
    display: flex;
    flex-direction: column;

}

#bio-button {
    border-radius: .2em;
    margin-left: .5em;
    border-color: whitesmoke;
    background-color: transparent;
    color: whitesmoke;
    border-style: double;
    margin-top: 1em;
    margin-bottom: 1em;

}

#bio-button:hover {
    cursor: pointer;
}

#figures {
    align-self: center;
    list-style-type: none;

}

.figure {
    border-top: .2em solid whitesmoke;
    margin: .1em;
}

.figure-name {
    font-size: 20pt;

}

#figure-image {
    max-width: 25%;
    max-height: 25%;
}

#timeline-title {
    margin-top: 1em;
    padding-top: 1em;
    border-top: .3em solid whitesmoke;
}

#search-bar {
    max-width: 30%;
    padding-top: 3em;
    align-self: center;
}

#search-input {
    border-radius: .5em;
    background-color: whitesmoke;
    border: none;
    outline-color: #cc683e;
}

#search-button {
    border-radius: .2em;
    margin-left: .5em;
    background-color: whitesmoke;
    border: none;

}

#search-button:hover {
    cursor: pointer;
}



.timeline-description {
    font-size: 12pt;
    max-width: 60%;
}
</style>