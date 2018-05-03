class UserController < ApplicationController
    def findUserById
        user = User.find_by(id: params[:id])
        render json: user
    end

    def findUserByUsername
        user = User.find_by(username: params[:username])
        render json: user
    end
end
