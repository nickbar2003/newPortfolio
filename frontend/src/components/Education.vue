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
      <h2 id="subheader"> Courses</h2>
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

      </ul>
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
  border: .05em solid whitesmoke;
  font-size: 15pt;
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
  padding: 1em;
  font-size: 20pt;
  list-style: circle;
}

.subbullet-description {
  font-size: 15pt;
  padding: .5em;
}
</style>
