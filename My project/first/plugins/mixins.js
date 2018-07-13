import Vue from 'Vue'

Vue.mixin({
    methods: {
      normal: function (oDate) {
      var iYear = oDate.getFullYear();
      var iMonth = oDate.getMounth();
      var iDay = oDate.getDate();
      return `${iYear}-${(iMonth >= 10 ? iMonth: `0${iMonth}`)}-${(iDay >= 10 ? iDay : `0${iDay}`)}`
      }
    }
  })