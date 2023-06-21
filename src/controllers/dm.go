package controllers

import (
	"gochat/src/controllers/authentication/authMiddleware"
	"gochat/src/controllers/authentication/userContext"
	"gochat/src/services/dm"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func AddDMController(myRouter *mux.Router) {
	// Get dms by workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm", authMiddleware.VerifyTokenMiddleware(getDMsByWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm", authMiddleware.VerifyTokenMiddleware(createDM)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm/{id}/join", authMiddleware.VerifyTokenMiddleware(joinToDM)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm/{id}/leave", authMiddleware.VerifyTokenMiddleware(leaveDM)).Methods("POST")

}

func getDMsByWorkspace(w http.ResponseWriter, r *http.Request) {
	
	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	dms, err, statusErr := dm.GetDMsByUserAndWorkspace(userContext, workspaceKey)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}

	w.Write(dms)
}

func createDM(w http.ResponseWriter, r *http.Request) {
	
	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	err, statusErr := dm.CreateDM(workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("DM created successfully"))
}

func joinToDM(w http.ResponseWriter, r *http.Request) {
	
	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err, statusErr := dm.JoinDM(id, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("DM joined successfully"))
}

func leaveDM(w http.ResponseWriter, r *http.Request) {
	
	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err, statusErr := dm.LeaveDM(id, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("DM left successfully"))
}
