After creating file and folder run below command in order to create a module

package > go mod init sample/project

Vs code -> code -> click preference -> settings -> search gopls -> edit replace the below items in that file

"gopls": {
        "experimentalWorkspaceModule": true
    }