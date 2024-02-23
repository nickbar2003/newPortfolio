
<script setup>
</script>

<template>
    <div id="work">
        <h2 id="header"> Work and Study </h2>
        <ul id="work-list">
            <li class="work-item"><strong>Groundskeeper</strong> <span class="date">Summer 2023</span></li>
            <li class="location">The Coeur d’Alene Gold & Spa Resort | Coeur d’Alene, ID</li>
            <ul class="work-description">

                <li class="work-description-bullet">Landscaping: mowing, trimming, edging.</li>
                <li class="work-description-bullet">Pool matinence: chemical testing, balance, cleaning.</li>
                <li class="work-description-bullet">Floral care: pruing, fertilizing, planting</li>
            </ul>

            <li class="work-item"><strong>Student and team member</strong> <span class="date">Summer 2022</span></li>
            <li class="location">Rennes Two University | Incoronata, Italy</li>
            <ul class="work-description">
                <li class="work-description-bullet">I had first hand experience doing the science of archaeology on an Early
                    Iron Age
                    Italian Site called <em>Incoranata</em></li>

                <li class="work-description-bullet">Excavation and recovery of Italian and Greek artifacts</li>
                <li class="work-description-bullet">Cleaning and idenfitfaction of artifacts</li>
                <li class="work-description-bullet">Stratigraphical documentation and analysis</li>
            </ul>
        </ul>

    </div>
</template>


<style scoped>
#work {
    margin-top: 10em;
    display: flex;
    flex-direction: column;
    border-top: .5em solid whitesmoke;
}


#header {
    font-size: 35pt;
}


#work-list {
    list-style: none;
}

.work-item {
    font-size: 20pt;
    padding-top: 1em;
    padding-left: 1em;
    padding-right: 1em;
    padding-bottom: 1em;
    display: flex;
    justify-content: space-between;
}

.work-description {
    padding: .5em;
    padding-left: 5em;
    font-size: 15pt;
    max-width: 60%;
}

.work-description-bullet {
    padding: .3em;
}

.location {
    padding-left: 1.5em;
    font-size: 15pt;

}

.date {
    font-size: 20pt;
}
</style>
