import axios from 'axios'

const API_VERSION = '/api/';

var service = {
    
    bookingUpsert: function (data, cancelCallback) {
        return axios.put(API_VERSION + 'booking', data);
    },

    bookingList: function (data,cancelCallback) {
        return axios.get(API_VERSION + 'booking/' + data);
    },

    eventList: function (cancelCallback) {
        return axios.get(API_VERSION + 'events');
    },

    eventUpsert: function (data, cancelCallback) {
        return axios.put(API_VERSION + 'events', data);
    },

    version: function (query, cancelCallback) {
        return axios.get(API_VERSION + 'version');
    }
}

export default service
