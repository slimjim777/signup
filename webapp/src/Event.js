import React, { Component } from 'react';
import moment from 'moment';
import DatePicker from 'react-datepicker';
import api from './api'

class Event extends Component {
    constructor(props) {
        
        super(props)

        var d = new Date();
        d.setDate(d.getDate() + (3 + 7 - d.getDay()) % 7)

        this.state = {
            id: null,
            name: '',
            date: moment(d),
            message: null,
            events: [],
        }

        this.getEvents()
    }

    getEvents() {
        api.eventList().then(response => {
            if (response.data.success) {
                this.setState({events: response.data.events})
            }
        })
    }

    validate() {
        var name = this.state.name.trim()
        if (name.length === 0) {
            this.setState({message: 'Enter the event name'})
            return false
        } else {
            this.setState({message: null})
            return true
        }
    }

    handleChangeEventDate = (date) => {
        this.setState({date: date, message: null})
    }

    handleChangeName = (e) => {
        this.setState({name: e.target.value, message: null});
    }

    handleSaveClick = (e) => {
        e.preventDefault()
        if (!this.validate()) {
            return
        }
        var b = {
            id: this.state.id,
            name: this.state.name,
            date: this.state.date.format('YYYY-MM-DD')
        }

        api.eventUpsert(b).then(response => {
            if (response.data.success) {
                this.getEvents()
                this.setState({id: null, name: '', message: null})
            }
        })
    }

    handleCancelClick = (e) => {
        e.preventDefault()
        this.setState({id: null, name: ''})
    }

    handleEditClick = (e) => {
        e.preventDefault()
        var eventId = parseInt(e.target.getAttribute('data-key'), 10)
        var event = this.state.events.filter((ev) => {
            return ev.id === eventId
        })[0]
        this.setState({
            id: event.id, name: event.name, date: moment(event.date),
            message: null
        })
    }

    renderAlert() {
        if (this.state.message) {
            return (
                <div className="alert">
                    {this.state.message}
                </div>
            )
        }
    }

    renderEvents() {
        return (
            <table className="event">
                {this.state.events.map(b => {
                    return (
                        <tr key={b.id}>
                            <td className="small">
                                <button className="event" onClick={this.handleEditClick} data-key={b.id}>Edit</button>
                                <button>Delete</button>
                            </td>
                            <td>{moment(b.date).format('LL')}</td>
                            <td>{b.name}</td>
                        </tr>
                    )
                })}
            </table>
        )
    }

    render() {
        return (
            <div className="row">
                <h1>Events</h1>
                {this.renderAlert()}
                <form>
                    <fieldset>
                        <input type="hidden" name="id" value={this.state.id} />
                        <label htmlFor="date">Date:
                            <DatePicker selected={this.state.date} onChange={this.handleChangeEventDate} dateFormat="LL" />
                        </label>
                        <label htmlFor="name">Event Name:
                            <input type="text" name="name" placeholder="name of the event or session"
                                    value={this.state.name} onChange={this.handleChangeName} />
                        </label>

                        <div className="clear">
                            <a href="" onClick={this.handleSaveClick} className="brand">Save</a>
                            &nbsp;
                            <button onClick={this.handleCancelClick}>Cancel</button>
                        </div>
                    </fieldset>
                </form>

                <div className="box clear">
                    <h3>Upcoming Events: {this.state.events.length}</h3>
                    <div className="row">
                        {this.renderEvents()}
                    </div>
                </div>

            </div>
        )
    }

}

export default Event;
