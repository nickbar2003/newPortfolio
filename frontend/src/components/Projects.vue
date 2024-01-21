<script setup>
</script>

<template>
  <div id="projects">
    <h2 id="header"> Projects </h2>
    <ul id="project-list">
      <li class="project-item"> Portfolio Website </li>
      <p class="project-description"> This webpage stores information about my past and present projects, my education
        information,
        current and past classes, as well as my work and study expriences.</p>
      <li class="project-item"> C++ Checkers</li>
      <p class="project-description">Utilizing C++ and the RayLib graphics library I programed a version of the popular
        game checkers.</p>
      <p class="project-description">
        You can check it out <a href="https://github.com/nickbar2003/C-Checkers"><strong>Here</strong></a>
      </p>
      <li class="project-item"> Historical Timeline </li>
      <p class="project-description">Combining my passion for history and computers this project utilizes <em>Vite and
          Vue</em>, <em>GoLang</em>, <em>JavaScript</em>, and <em>Postgres</em> to allow a user to search up a historical
        figure to see many other important figures that were alive at the same time.</p>
      <p class="project-description">
        You can check it out <router-link id="here" to='/Timeline'><strong>Here.</strong></router-link>
      </p>
      <li class="project-item"> Accurate Roman Name Generator </li>
      <p class="project-description">This is a terminal based application I made using C++ in my
        spare time while taking CS 121. Based on the day of the week and the liturgical season the program will pray the
        correct rosary. I made this to implement some of the concepts I had learned so far</p>
    </ul>
  </div>
</template>


<style scoped>
#projects {
  margin-top: 10em;
  display: flex;
  flex-direction: column;
  border-top: .5em solid whitesmoke;
}


#header {
  font-size: 35pt;
}

#here {
  color: purple;
  text-decoration: none;
  padding: .2em;
  transition: all 400ms ease-in-out;

}

#here:hover {
  color: purple;
  text-decoration: none;
  background-color: whitesmoke;
}

#project-list {
  list-style: none;
}

.project-item {
  font-size: 20pt;
  padding: 1em;
}

.project-description {
  padding-left: 5em;
  font-size: 12pt;
  max-width: 60%;
}
</style>
