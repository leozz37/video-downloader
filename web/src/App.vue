<template>
  <div class="govideo">
    <h1>Welcome to Go-Video</h1>
    <p> Paste the video URL into the text box:</p>

    <form v-on:submit.prevent="postreq">
      <div class="form-group">
        <input v-model="url" type="url" id="url-input" name="user_input">
      </div>
    </form>

    <div class="instagram">
      <h3>Supported plataforms</h3>
      <a href="https://youtube.com" target="_blank" class="noopener"><img src="./assets/youtube.png"></a>
      <a href="https://twitter.com" target="_blank" class="noopener"><img src="./assets/twitter.png"></a>
      <a href="https://instagram.com" target="_blank" class="noopener"><img src="./assets/instagram.png"></a>
      <a href="https://facebook.com" target="_blank" class="noopener"><img src="./assets/facebook.png"></a>
      <a href="https://twitch.com" target="_blank" class="noopener"><img src="./assets/twitch.png"></a>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'govideo',

  data() { return {
    url: '',
  } },

  methods: {
    postreq() {
      axios.post("http://127.0.0.1:8090/download", {
      data: this.url,
      responseType: 'blob'
      })
        .then(response => {
            var blob = new Blob([response.data]);
            var downloadElement = document.createElement("a");
            var href = window.URL.createObjectURL(blob);
            downloadElement.href = href;
            downloadElement.download = "test.mp4";
            document.body.appendChild(downloadElement);
            downloadElement.click();
            document.body.removeChild(downloadElement);
            window.URL.revokeObjectURL(href);
          })
        .catch(response => {
          console.log(response);
        });
    }, 
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.govideo {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
h3 {
  margin: 40px 0 0;
}
a {
  color: #42b983;
}
img {
  width: 42px;
  height: 42px;
  margin: 10px;

}
input {
  width: 300px;
  height: 20px;
  border: 0;
  border-bottom: 2px solid gray;
  outline: 0;
  background: transparent;
  transition: border-color 0.2s;
  margin: 0 10px;
}
</style>