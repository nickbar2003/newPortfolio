<template>
  <div id="app">
    <header>
      <div>
        <nav>
          <router-link to="/" class="link">Home</router-link>
          <router-link to="/Portfolio" class="link">Portfolio</router-link>
          <router-link to="/Timeline" class="link">Timeline</router-link>
        </nav>
      </div>
    </header>


    <main>
      <router-view></router-view>
    </main>



  </div>
</template>

<style scoped>
#app {}




nav {
  display: flex;
  justify-content: space-evenly;
  font-size: 2em;
}

.link {
  display: flex;
  color: whitesmoke;
  background-color: inherit;
  text-decoration: none;
  border-radius: .1em;
  padding: .5em;
  transition: all 400ms ease-in-out;
}

.link:hover {

  color: #cc683e;
  background-color: whitesmoke;
  transform: scale(1.2, 1.2);
}
</style>
