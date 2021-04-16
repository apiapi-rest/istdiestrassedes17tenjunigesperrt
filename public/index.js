window.onload = () => {
  fetch(
    window.location.host === `localhost:8080`
      ? `http://localhost:8081/availability`
      : `https://europe-west3-istdiestrassedes17tenjunigespe.cloudfunctions.net/availability`,
    {
      method: "GET",
      cors: false,
      headers: {},
    }
  )
    .then((result) => result.json())
    .then((result) => {
      console.log(result);

      if (result.blocked) {
        document.querySelector("#ja").classList.add("show");
      } else {
        document.querySelector("#nein").classList.add("show");
      }
    });
};
