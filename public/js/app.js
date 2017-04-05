var app = new Vue({
  el: '#app',
  data: {
    buzzword: ''
  },
  created: function () {
    this.getBuzzword();
  },
  methods: {
    getBuzzword: function () {
        var self = this;
        $.get("/buzzword", function(data) {
            self.buzzword = data;
        })
    },
    getBuzzwordWithSuffix: function () {
        var self = this;
        $.get("/suffix", function(data) {
            self.buzzword = data;
        })
    },
    getBuzzwordWithVerb: function () {
        var self = this;
        $.get("/verb", function(data) {
            self.buzzword = data;
        })
    },
    getBuzzwordWithVerbAndSuffix: function () {
        var self = this;
        $.get("/verbsuffix", function(data) {
            self.buzzword = data;
        })
    }
  }

})