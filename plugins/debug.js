import Vue from 'vue'
import mitt from '@stephendltg/e-bus'

const emitter = mitt()

const { log, clear, group, groupEnd, groupCollapsed } = console

// Clear console
clear()

const size = (value) => {
  if (value === null) {
    return null
  } else if (typeof value === 'number') {
    return value
  } else if (typeof value === 'string') {
    return value.trim().length
  } else if (Array.isArray(value)) {
    return value.length
  } else if (typeof value === 'object') {
    return value.size || Object.keys(value).length
  } else {
    return null
  }
}

const color = mitt()

const print = (name, val) => {
  name = name.toString()
  if (process.env.debug === '*' || process.env.debug === name) {
    if (!color.get(name)) {
      color.set(name, '#55' + name.length)
    }
    groupCollapsed('%cDebug: ' + name + ' - [' + new Date(Date.now()).toISOString() + ']', 'color:white;background-color:' + color.get(name) + ';')
    log('%c data: ', 'color:white;background-color:black;', val)
    log('%c type: ', 'color:white;background-color:blue;', typeof val)
    log('%c length: ', 'color:white;background-color:orange;', size(val))
    group('Output')
    if (val && typeof val !== 'function' && typeof val !== 'symbol') {
      log(JSON.parse(JSON.stringify(val)))
    }
    groupEnd()
    log('%c', 'color:white;background-color:orange;')
    groupEnd()
  }
}

const debug = (val, name = '*') => process.env.dev ? print(name, val) : null

export default ({ app }, inject) => {
  inject('mitt', Vue.observable(emitter))
  inject('debug', Vue.observable(debug))
}
