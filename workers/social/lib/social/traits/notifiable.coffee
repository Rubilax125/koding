KodingError = require '../error'

module.exports = class Notifiable

  updateAndNotify: (options, change, callback) ->

    { account, group, target } = options

    id = @getId()

    @update change, (err) ->
      callback err

      switch target
        when 'group'

          JGroup = require '../models/group'
          JGroup.one { slug : group }, (err, group_) ->
            return  if err or not group_

            opts = { id, group, change, timestamp: Date.now() }
            group_.sendNotification 'InstanceChanged', opts, ->

        when 'account'

          opts = { id, group, change, timestamp: Date.now() }
          account.sendNotification 'InstanceChanged', opts

