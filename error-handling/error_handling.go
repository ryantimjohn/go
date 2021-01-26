package erratum

// Use returns an error
func Use(o ResourceOpener, input string) (e error){
   var r Resource

   for r, e = o(); e != nil; r, e = o(){
     if _, ok := e.(TransientError); !ok{
       return e
     }
   }

   defer r.Close()
   defer func(){
     if rec := recover(); rec != nil {
       _, ok := rec.(FrobError)
       if ok == true{
         frobErr := rec.(FrobError)
         r.Defrob(frobErr.defrobTag)
       }
       e = rec.(error)
     }
   }()

   r.Frob(input)

   return e
}