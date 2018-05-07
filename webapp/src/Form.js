import React, { Component } from 'react';
import moment from 'moment';
import api from './api'

class Form extends Component {

    constructor(props) {
        
        super(props)

        var d = new Date();
        d.setDate(d.getDate() + (3 + 7 - d.getDay()) % 7);

        this.state = {
            name: localStorage.getItem('name') || '',
            message: null,
            monday: moment(d),
            bookings: [],
        }

        this.getBookings()
    }

    getBookings() {
        api.bookingList(this.state.monday.format('YYYY-MM-DD')).then(response => {
            if (response.data.success) {
                this.setState({bookings: response.data.bookings})
            }
        })
    }

    validate() {
        var name = this.state.name.trim()
        if (name.length === 0) {
            this.setState({message: 'Dude! Enter your name!'})
            return false
        } else {
            this.setState({message: null})
            return true
        }
    }

    arePlaying() {
        return this.state.bookings.filter(b => {
            return b.playing;
        })
    }

    handleChangeName = (e) => {
        this.setState({name: e.target.value, message: null});
        if (e.target.value !== '') {
            localStorage.setItem('name', e.target.value)
        }
    }

    handleYesClick = (e) => {
        e.preventDefault()
        if (!this.validate()) {
            return
        }

        var b = {
            name: this.state.name,
            date: this.state.monday.format('YYYY-MM-DD'),
            playing: true
        }
        api.bookingUpsert(b).then(response => {
            this.getBookings()
        })
    }

    handleNoClick = (e) => {
        e.preventDefault()

        if (!this.validate()) {
            return
        }

        var b = {
            name: this.state.name,
            date: this.state.monday.format('YYYY-MM-DD'),
            playing: false
        }
        api.bookingUpsert(b).then(response => {
            this.getBookings()
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

    renderBookings() {
        var booked = this.arePlaying()
    
        return (
            <table>
                {booked.map(b => {
                    return (
                        <tr key={b.id}>
                            <td>{b.name}</td>
                        </tr>
                    )
                })}
            </table>
        )
    }

    render () {

        return (
            <div className="row">
                {this.renderAlert()}
                <form>
                    <fieldset>
                        <label htmlFor="monday">Date:
                            <b>
                            <input type="text" disabled
                                    value={this.state.monday.format('DD/MM/YYYY')} />
                            </b>
                        </label>
                        <label htmlFor="name">Name:
                            <input type="text" placeholder="your name, anointed one"
                                    value={this.state.name} onChange={this.handleChangeName} />
                        </label>

                        <div className="clear">
                            <a href="" onClick={this.handleYesClick} className="brand">Attending</a>
                            &nbsp;
                            <button onClick={this.handleNoClick}>Not Attending</button>
                        </div>
                    </fieldset>
                </form>

                <div className="box clear">
                    <h3>Signed Up: {this.arePlaying().length}</h3>
                    <div className="row">
                        {this.renderBookings()}
                    </div>
                </div>
            </div>
        )
    }
}

export default Form;
