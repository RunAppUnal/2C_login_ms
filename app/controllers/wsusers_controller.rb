class WsusersController < ApplicationController
    soap_service namespace: 'urn:WashOutUser', camelize_wsdl: :lower

    # check account case
    soap_action "getUser",
                :args   => { :username => :string },
                :return => { :userValid => :boolean, :name => :string, :lastname => :string, :email => :string, :cellphone => :string}
    def getUser
      if (User.exists?(username: params[:username]))
          @user = User.find_by(username: params[:username])
          name = @user.name
          lastname = @user.lastname
          email = @user.email
          cellphone = @user.cellphone
          render :soap => { :userValid => true, :name => name, :lastname => lastname, :email => email, :cellphone => cellphone }
      else
          render :soap => { :userValid => false, :name => "", :lastname => "", :email => "", :cellphone => ""}
      end
    end
  end
