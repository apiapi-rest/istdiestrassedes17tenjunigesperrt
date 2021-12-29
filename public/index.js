window.onload = () => {
  fetch(
    window.location.host === `localhost:8080`
      ? `http://localhost:3000/availability`
      : `https://istdiestrassedes17tenjunigesperrt.apiapi.rest/availability`,
    {
      method: "GET",
      cors: false,
      headers: {},
    }
  )
    .then((result) => result.json())
    .then((result) => {
      console.log(result);

      if (result.success === true && result.data.blocked) {
        document.querySelector("#ja").classList.add("show");
      } else {
        document.querySelector("#nein").classList.add("show");
      }
    });
};
