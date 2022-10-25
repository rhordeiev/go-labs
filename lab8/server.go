package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ActiveLink struct {
	Active int
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("website/js"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("website/img/"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("website/css/"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("website/fonts/"))))
	r.PathPrefix("/color/").Handler(http.StripPrefix("/color/", http.FileServer(http.Dir("website/color/"))))
	r.PathPrefix("/residency/js/").Handler(http.StripPrefix("/residency/js/", http.FileServer(http.Dir("website/js"))))
	r.PathPrefix("/residency/img/").Handler(http.StripPrefix("/residency/img/", http.FileServer(http.Dir("website/img/"))))
	r.PathPrefix("/residency/css/").Handler(http.StripPrefix("/residency/css/", http.FileServer(http.Dir("website/css/"))))
	r.PathPrefix("/residency/fonts/").Handler(http.StripPrefix("/residency/fonts/", http.FileServer(http.Dir("website/fonts/"))))
	r.PathPrefix("/residency/color/").Handler(http.StripPrefix("/residency/color/", http.FileServer(http.Dir("website/color/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ActiveLink{
			Active: 1,
		}
		tmpl, _ := template.ParseFiles("./website/index.html", "./website/header.html", "./website/footer.html")
		tmpl.Execute(w, data)
	})
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := ActiveLink{
			Active: 2,
		}
		tmpl, _ := template.ParseFiles("./website/about.html", "./website/header.html", "./website/footer.html")
		tmpl.Execute(w, data)
	})
	r.HandleFunc("/residency", func(w http.ResponseWriter, r *http.Request) {
		type ActiveLinkResidency struct {
			Active  int
			Numbers []int
		}
		data := ActiveLinkResidency{
			Active:  3,
			Numbers: []int{1, 2, 3, 4, 5, 6},
		}
		tmpl, _ := template.ParseFiles("./website/living.html", "./website/header.html", "./website/footer.html")
		tmpl.Execute(w, data)
	})
	r.HandleFunc("/residency/{id:[1-6]}", func(w http.ResponseWriter, wr *http.Request) {
		type Room struct {
			Name        string
			Description string
		}
		type ActiveLinkResidencyRoom struct {
			Active int
			Room   Room
			Id     int64
		}
		Rooms := []Room{
			{"НОМЕР ДЕЛЮКС", `Номери делюкс - це плавні сучасні лінії і закруглені кути, що нагадують морські хвилі, 
			накочуються на берег під приватним балконом вашого номера.
			Колірна палітра інтер'єру включає в себе відтінки приглушеного синього, зеленого і сірувато-білого,
			 створюючи атмосферу спокою. У розпорядженні гостей, що зупинилися в номері, знаходиться велике двоспальне ліжко,
			  оригінальні приставні столики, туалетний столик, кушетка і сучасний Smart-телевізор (з можливістю підключення персональних мобільних пристроїв).
			   Номер делюкс площею 45 м² включає простору ванну кімнату з мармуровою обробкою, подвійними раковинами, великою ванною і окремою душовою.
			
			Номери делюкс призначені не більше ніж для трьох гостей і ідеальні для пар та молодих сімей.`},
			{"НОМЕР ЛЮКС З ДВОМА СПАЛЬНЯМИ", `Відмінні риси номерів люкс - дубовий паркет і простора ванна кімната, що забезпечують найвищий комфорт і неперевершену зручність.
			Білене дерево, витончені природні тони і шикарні м'які меблі створюють чудову атмосферу. У номері є велике двоспальне ліжко, кушетка і сучасний Smart-телевізор
			 (з можливістю підключення персональних мобільних пристроїв), а також автоматична декоративне підсвічування.
			
			У номері площею 135 м² можуть з комфортом розміститися двоє дорослих і двоє дітей або пари, які бажають насолодитися розкішним відпочинком.`},
			{"НОМЕР ЛЮКС З ОДНІЄЮ СПАЛЬНЕЮ", `Відмінні риси номерів люкс - дубовий паркет і простора ванна кімната, що забезпечують найвищий комфорт і неперевершену зручність.
			Білене дерево, витончені природні тони і шикарна м'які меблі створюють чудову атмосферу. У номері є велике двоспальне ліжко, кушетка і сучасний Smart-телевізор 
			(з можливістю підключення персональних мобільних пристроїв), а також автоматична декоративне підсвічування.
			
			Номер люкс з однією спальнею - це ідеальний вибір для тих, хто любить подорожувати всією сім'єю, так як на площі 90 м² можуть комфортно розміститися
			 двоє дорослих і двоє дітей.`},
			{"НАПІВЛЮКС", `Напівлюкси готелю AMARA відрізняються чарівною простотою і сучасним стилем.
			 Напівлюкс складається з однієї спальні з примикає до неї ванною кімнатою, окремої вітальні і обідньої зони, а також окремою ванною кімнати і душовою.
			  Цей номер відмінно підходить для комфортного відпочинку далеко від повсякденної суєти. Розташуєтеся на м'якому дивані,
			на широкому балконі з лежаками або на великим двоспальним ліжка, і відчуйте себе як вдома в оточенні затишного дерев'яного декору і ніжних пастельних тонів.
			 
			 У напівлюксі площею 80 м² можуть комфортно розміститися до чотирьох гостей (двоє дорослих і двоє дітей).`},
			{"ДВОРІВНЕВИЙ ЛЮКС НА ДАХУ З ПРИВАТНИМ БАСЕЙНОМ", `Цей ексклюзивний дворівневий люкс, призначений тільки для дорослих, є відмінним вибором для тих,
			  хто вважає за краще спокій і самота, і втіленням справжньої краси.
			 Дворівневий люкс оснащений панорамними вікнами, що забезпечують ідеальне природне освітлення. Дизайн відрізняється плавними лініями і мінімалістичністю
			  з переважанням вершкових і кремових тонів. Інтер'єр поєднує в собі сучасну елегантність і чарівність класики. Площа дворівневого люкса на даху становить 110 м²,
			   і в розпорядженні гостей є простора ванна кімната з подвійними раковинами, великою ванною і окремою душовою, а також затишна зона відпочинку з Smart-телевізором
			    (з можливістю підключення персональних мобільних пристроїв). Приватний басейн в саду з чудовими видами - відмінне доповнення для пар,
				 що знаходяться в пошуках виняткової розкоші.`},
			{"КОТЕДЖ З ВИДОМ НА МОРЕ З ОДНІЄЇ СПАЛЬНЕЮ, ВІТАЛЬНЕЮ ТА ВЛАСНИМ БАСЕЙНОМ", `Котедж з однією спальнею відрізняється прекрасним видом на море і володіє
				  всіма перевагами традиційного кіпрського стилю: від широких вікон, впускає всередину промені сонячного світла, до предметів інтер'єру з текстурою вугілля і каменю,
				   відмінно поєднуються з місцевими матеріалами.
				 Цей розкішний номер люкс площею 125 м² з власним басейном призначений для двох дорослих і двох дітей і ідеально підходить як для пар,
				  які вирушають у свій довгоочікуваний відпустку, так і для сімей, які шукають відмінне місце для відпочинку і дослідження узбережжя Лімасола.`},
		}
		vars := mux.Vars(wr)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		data := ActiveLinkResidencyRoom{
			Active: 3,
			Room:   Rooms[id-1],
			Id:     id,
		}
		tmpl, _ := template.ParseFiles("./website/room_description.html", "./website/header.html", "./website/footer.html")
		tmpl.Execute(w, data)
	})
	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		data := ActiveLink{
			Active: 4,
		}
		tmpl, _ := template.ParseFiles("./website/contact.html", "./website/header.html", "./website/footer.html")
		tmpl.Execute(w, data)
	})

	http.Handle("/", r)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)

}
