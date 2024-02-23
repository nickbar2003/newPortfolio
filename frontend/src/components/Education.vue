<script setup></script>

<template>
  <div id="education">
    <h2 id="header"> Education | Gonzaga University</h2>
    <ul id="majors-list">
      <li class="major-name">Bachelor of Arts, Computer Science & Computational Thinking (Exp Comp 5/25)</li>
      <li class="major-name">Bachelor of Arts, History</li>
      <li class="major-name">Minor in Classical Civilizations</li>
    </ul>
    <div id="gpa"> Cumulative GPA: 3.94 | President’s list: Four Semesters</div>
    <div class="courses">
      <h2 id="subheader">Courses</h2>
      <ul id="courses-list">
        <li class="course-name">CPSC 260 Computer Organization
          <ul class="course-description">
            <li class="course-description-bullet">Digital Logic and Boolean Algebra</li>
            <li class="course-description-bullet">C programming</li>
            <li class="course-description-bullet">Circuit synthesis and analysis</li>
            <li class="course-description-bullet">Computer Architecture</li>
            <li class="course-description-bullet">Assembly Language</li>
            <li class="course-description-bullet">Caching and Basic Code Optimization Approaches</li>
          </ul>
        </li>
        <li class="course-name">CPSC 222 Introduction to Data Science
          <ul class="course-description">
            <li class="course-description-bullet">Data Science Lifecycle</li>
            <li class="course-description-bullet">Data gathering and cleaning</li>
            <li class="course-description-bullet">Data visualization</li>
            <li class="course-description-bullet">Model formulation</li>
            <li class="course-description-bullet">Utilizes Python and data science packages</li>
          </ul>
        </li>
        <li class="course-name">CPSC 331 UI/UX Design
          <ul class="course-description">
            <li class="course-description-bullet">Human-centered software design processes</li>
            <li class="course-description-bullet">Data gathering methods</li>
            <li class="course-description-bullet">Low and high-fidelity prototyping</li>
            <li class="course-description-bullet">Empirical and Analytical software analysis</li>
          </ul>
        </li>
        <li class="course-name">CPSC 223 Algorithm Abstract Data Struct
          <ul class="course-description">
            <li class="course-description-bullet">Data Structures</li>
            <li class="course-description-bullet">Algorithm design and implementation, including sorting algorithms</li>
            <li class="course-description-bullet">Asymptotic analysis of algorithms</li>
            <li class="course-description-bullet">Advanced C++ concepts</li>
          </ul>
        </li>
        <li class="course-name">CPSC 212 Computational Modeling
          <ul class="course-description">
            <li class="course-description-bullet">Problems of organized complexity</li>
            <li class="course-description-bullet">Identify and explain emergent behaviors that arise from agent
              interactions</li>
            <li class="course-description-bullet">Formulate Agent Based Models</li>
            <li class="course-description-bullet">Analyze and interpret patterns in simulated systems</li>
            <li class="course-description-bullet">Validate and verify models, including sensitivity analysis</li>
          </ul>
        </li>
        <li class="course-name">CPSC 121 Computer Science I
          <ul class="course-description">
            <li class="course-description-bullet">Introduction to Computer Science with C++</li>
            <li class="course-description-bullet">Algorithm design</li>
            <li class="course-description-bullet">Program design and implementation</li>
            <li class="course-description-bullet">Introduction to Linux and Bash</li>
            <li class="course-description-bullet">Git and Github</li>
          </ul>
        </li>
        <li class="course-name">CPSC 122 Computer Science II
          <ul class="course-description">
            <li class="course-description-bullet">Object-oriented programming</li>
            <li class="course-description-bullet">Sorting and searching algorithms</li>
            <li class="course-description-bullet">Linear data structures, including linked lists, stacks, and queues</li>
          </ul>
        </li>
        <li class="course-name">CPSC 224 Software Development
          <ul class="course-description">
            <li class="course-description-bullet">Identify, specify, and analyze functional requirements</li>
            <li class="course-description-bullet">Design and implement an application with a graphical user interface</li>
            <li class="course-description-bullet">Collaborate with others on a team project utilizing Git and GitHub</li>
          </ul>
        </li>
        <li class="course-name">CPSC 332 Web Development
          <ul class="course-description">
            <li class="course-description-bullet">Techniques of web page development</li>
            <li class="course-description-bullet">HTML, CSS, JavaScript, EJS</li>
            <li class="course-description-bullet">DOM Structure and manipulation</li>
            <li class="course-description-bullet">Full stack development</li>
            <li class="course-description-bullet">Node.js</li>
            <li class="course-description-bullet">RESTful web API’s</li>
            <li class="course-description-bullet">MySQL database</li>
          </ul>
        </li>
      </ul>

      <!-- <h2 id="subheader"> Courses</h2>
      <ul id="courses-list">
        <li class="course-name"> CPSC 224 </li>
        <ul class="course-description">
          <li class="course-description-bullet"> Continuation of software development </li>
          <ul class="course-description-subbullet">
            <li class="subbullet-description"> Identification of program requirements</li>
            <li class="subbullet-description"> Program design and implementation with graphical user interface</li>
            <li class="subbullet-description"> Git and GitHub in the context of team collaboration</li>
            <li class="subbullet-description"> Unit testing </li>
          </ul>
          <li class="course-description-bullet"> object-oriented programming with Java </li>
        </ul>
        <li class="course-name"> CS 322 </li>
        <ul class="course-description">
          <li class="course-description-bullet"> Continuation of Programming with C++ </li>
          <li class="course-description-bullet"> Learned more concepts of software development </li>
          <ul class="course-description-subbullet">
            <li class="subbullet-description"> More complex data structures </li>
            <li class="subbullet-description"> Continue program design and implementation</li>
            <li class="subbullet-description"> Algorithm efficiency analysis</li>
            <li class="subbullet-description"> Git and Github for software version control</li>
          </ul>
        </ul>

        <li class="course-name"> CPSC 121 </li>
        <ul class="course-description">
          <li class="course-description-bullet"> Introduction to Computer Science with C++</li>
          <ul class="course-description-subbullet">
            <li class="subbullet-description"> Algorithm design</li>
            <li class="subbullet-description"> Program design and implementation</li>
            <li class="subbullet-description"> Git and GitHub</li>
            <li class="subbullet-description"> Data structure design and implementation</li>
          </ul>

        </ul>

      </ul> -->

    </div>
  </div>
</template>


<style scoped>
#header {
  font-size: 35pt;
}

#education {
  margin-top: 5em;
  border-top: .5em solid whitesmoke;
}

#majors {
  list-style-type: square;
}

.major-name {
  padding: 1em;
  font-size: 25pt;
}

#gpa {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  padding: 1em;
  border: .08em solid whitesmoke;
  font-size: 20pt;
  max-width: 30em;
}

#subheader {
  margin-top: 2em;
  border-top: .1em solid whitesmoke;
  padding: .5em;
  font-size: 25pt;
  max-width: 50%;
}

.course-name {
  padding: .5em;
  font-size: 25pt;
  list-style: none;
}

.course-description-bullet {
  padding: .5em;
  font-size: 20pt;
  list-style: circle;
}

.subbullet-description {
  font-size: 15pt;
  padding: .5em;
}

.course-description {

  margin-bottom: 3em;
}
</style>
