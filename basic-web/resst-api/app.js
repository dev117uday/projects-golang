var app = new Vue({
    el: '#app',
    data: {
      name: 'Uday Yadav'
    },
    methods : {
      getData : function() {
        fetch("https://jsonplaceholder.typicode.com/todos/1")
        .then(data => {
          return data.json()
        }).then(response => {
          console.log(response)
        }).catch(error => {
          console.log(error);
        })
      }
    }
  })