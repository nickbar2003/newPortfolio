<template>
    <div>
        <div id="header">
            <h1 id="intro"> Hi, I'm Nick! </h1>
            <img id="photo" src="/picture.jpg" alt="photo" />
            <h1 id="welcome"> Welcome To My Website </h1>
        </div>
    </div>

    <div id="secret-code">
        {{ secretMessage }}
    </div>
</template>

<script>
export default {
    data() {
        return {
            secretMessage: ""
        }
    },
    async mounted() {
        const resp = await fetch("http://localhost:3000/hello", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }

        })
        console.log(resp)
        const foo = await resp.json();
        console.log(foo)
    }

}
</script>


<style scoped>
#intro {
    font-size: 50pt;
    text-shadow: .05em .05em .5em whitesmoke;
}

#welcome {
    font-size: 35pt;
    text-shadow: .05em .05em .5em whitesmoke;
}

#photo {
    max-width: 35em;
    max-height: 35em;
    border-radius: 25%;
    border: .5em solid whitesmoke;
    transition: all 600ms ease-in-out;
    box-shadow: .05em .05em .5em whitesmoke;


}

#photo:hover {
    max-width: 40em;
    max-height: 40em;
    box-shadow: 0em 0em 5em;
    border: .3em solid whitesmoke;
}

#header {
    display: flex;
    flex-direction: column;
    align-items: center;

}
</style>
