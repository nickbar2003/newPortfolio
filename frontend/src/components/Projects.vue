<script setup>
</script>

<template>
  <div id="projects">
    <h2 id="header"> Projects </h2>
    <ul id="project-list">
      <li class="project-item"> Portfolio Website </li>
      <p class="project-description"> This Web page stores information about my past and present projects, my education
        information, current and past classes, as well as my work and study experiences.</p>
      <li class="project-item"> C++ Checkers</li>
      <p class="project-description">Utilizing C++ and the RayLib graphics library I programed a version of the popular
        game checkers.</p>
      <p class="project-description">
        You can check it out <a href="https://github.com/nickbar2003/C-Checkers"><strong>Here</strong></a>
      </p>
      <li class="project-item"> Historical Figure Search </li>
      <p class="project-description">Combining my passion for history and computers this project utilizes <em>Vite and
          Vue</em>, <em>Golang</em>, <em>JavaScript</em>, and <em>SQL</em> and <a
          href="https://github.com/trietmn/go-wiki?tab=readme-ov-file#2-getpage-there-are-multiple-page-methods-in-the-wikipedia-page-methods"><em>go-wiki</em></a>(a
        Golang Wikipedia API wrapper) to allow users to look up historical figures and discover other figures who
        were alive at the same time.</p>
      <p class="project-description">
        You can check it out <router-link id="here" to='/Timeline'><strong>Here.</strong></router-link>
      </p>
      <li class="project-item"> Accurate Roman Name Generator </li>
      <p class="project-description">This is a c++ program for generating historically plausible Roman names. With a
        Praenomen (first name), Gensnomen (clan name), and Cognomen (family name). The program can generate over 2 million
        names. I hope to expand this project to include Roman names for places like villages, cities, and coutnreis as
        well as geographic names for continents, regions, and bodies of water.</p>
      <p class="project-description">
        You can check it out <a href="https://github.com/nickbar2003/Roman-Name-Generator"><strong>Here</strong></a>
      </p>
    </ul>
  </div>
</template>


<style scoped>
#projects {
  margin-top: 10em;
  display: flex;
  flex-direction: column;
  /* align-items: center; */
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
  /* align-items: center; */
  display: flex;
  flex-direction: column;
}

.project-item {
  font-size: 20pt;
  padding: 1em;
}

.project-description {
  padding-left: 5em;
  font-size: 15pt;
  max-width: 60%;
}
</style>
