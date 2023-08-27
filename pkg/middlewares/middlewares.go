package middlewares

import (
    "net/http"
//    "fmt"
//    "context"

//    "github.com/go-chi/chi/v5"
//    "github.com/go-chi/chi/v5/middleware"
)

func ChangeMethod(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            switch method := r.PostFormValue("_method"); method {
            case http.MethodPut:
                fallthrough
            case http.MethodPatch:
                fallthrough
            case http.MethodDelete:
                r.Method = method
            default:
            }
        }
        next.ServeHTTP(w, r)
    })
}

func ArticleCtx(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//        articleID := chi.URLParam(r, "articleID")
//        article, err := dbGetArticle(articleID) // TODO: change to get chatroom ID
//        if err != nil {
//            fmt.Println(err)
//            http.Error(w, http.StatusText(404), 404)
//            return
//        }
//        ctx := context.WithValue(r.Context(), "article", article)
//        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
