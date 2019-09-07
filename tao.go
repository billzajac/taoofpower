package main

// FIXME: Finish adding passages 16-81
// Remember to rename the package name to tao and comment main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", Tao)
	http.ListenAndServe(":8080", nil)
}

func Tao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, GetPassage())
}

func GetPassage() string {
	var passages [81]string
	passages[0] = `THE BEGINNING OF POWER\n
The Tao that can be expressed\n
is not the Tao of the Absolute. The name that can be named\n
Is not the name of the Absolute.\n
The nameless originated Heaven and Earth. The named is the Mother of All Things.\n
Thus, without expectation,\n
One will always perceive the subtlety; And, with expectation,\n
One will always perceive the boundary.\n
The source of these two is identical, Yet their names are different. Together they are called profound, Profound and mysterious, The gateway to the Collective Subtlety.
`
	passages[1] = `USING POLARITY\n
When all the world knows beauty as beauty,\n
There is ugliness. When they know good as good,\n
There there is evil.\n
In this way\n
Existence and nonexistence produce each other. Difficult and easy complete each other. Long and short contrast each other. High and low attract each other. Pitch and tone harmonize each other. Future and past follow each other.\n
Therefore, Evolved Individuals\n
Hold their position without effort, Practice their philosophy without words, Are a part of All Things and overlook nothing. They produce but do not possess, Act without expectation, Succeed without taking credit.\n
Since, indeed, they take no credit, it remains with them.\n
`
	passages[2] = `KEEPING PEACE\n
Do not exalt the very gifted,\n
And people will not contend. Do not treasure goods that are hard to get,\n
And people will not become thieves. Do not focus on desires,\n
And people's minds will not be confused.\n
Therefore, Evolved Individuals lead oti is by\n
Opening their minds, Reinforcing their centers, Relaxing their desires, Strengthening their characters.\n
Let the people always act without strategy or desire;\n
Let the clever not venture to act. Act without action,\n
And nothing is without order.\n
`
	passages[3] = `THE NATURE OF THE TAO\n
The Tao is empty and yet useful; Somehow it never fills up. So profound! It resembles the source of All Things.\n
It blunts the sharpness, Unties the tangles, And harmonizes the brightness. It identifies with the ways of the world.\n
So deep! It resembles a certain existence. I do not know whose offspring it is, This Image in front of the source.\n
`
	passages[4] = `HOLDING TO THE CENTER\n
Heaven and Earth are impartial;\n
They regard All Things as straw dogs. Evolved Individuals are impartial;\n
They regard all people as straw dogs.\n
Between Heaven and Earth,\n
The space is like a bellows. The shape changes,\n
But not the form. The more it moves,\n
The more it produces.\n
Too much talk will exhaust itself. It is better to remain centered.\n
`
	passages[5] = `PERCEIVING THE SUBTLE\n
The mystery of the valley is immortal;\n
It is known as the Subtle Female. The gateway of the Subtle Female\n
Is the source of Heaven and Earth.\n
Everlasting, endless, it appears to exist. Its usefulness comes with no effort.\n
`
	passages[6] = `THE POWER OF SELFLESSNESS\n
Heaven is eternal, the Earth everlasting. They can be eternal and everlasting Because they do not exist for themselves. For that reason they can exist eternally.\n
Therefore, Evolved Individuals\n
Put themselves last,\n
And yet they are first. Put themselves outside,\n
And yet they remain.\n
Is it not because they are without self-interest That their interests succeed?\n
`
	passages[7] = `NONCOMPETITIVE VALUES\n
The highest value is like water.\n
The value in water benefits All Things,\n
And yet it does not contend. It stays in places that others disdain,\n
And therefore is close to the Tao.\n
The value in a dwelling is location. The value in a mind is depth. The value in relations is benevolence. The value in words is sincerity. The value in leadership is order. The value in work is competence. The value in effort is timeliness.\n
Since, indeed, they do not contend, There is no resentment.\n
`
	passages[8] = `TRANSCENDING DECLINE\n
Holding to fullness Is not as good as stopping in time.\n
Sharpness that probes Cannot protect for long.\n
A house filled with riches Cannot be defended.\n
Pride in wealth and position Is overlooking one's collapse.\n
Withdrawing when success is achieved\n
Is the Tao in Nature.\n
`
	passages[9] = `INNER HARMONY\n
In managing your instincts and embracing Oneness,\n
Can you be undivided? In focusing your Influence,\n
Can you yield as a newborn child? In clearing your insight,\n
Can you become free of error? In loving people and leading the organization,\n
Can you take no action? In opening and closing the gateway to nature,\n
Can you not weaken? In seeing clearly in all directions,\n
Can you be without knowledge?\n
\n
Produce things, cultivate things;\n
Produce but do not possess.\n
Act without expectation.\n
Advance without dominating.\n
These are called the Subtle Powers.\n
`
	passages[10] = `USING WHAT IS NOT\n
Thirty spokes converge at one hub;\n
What is not there makes the wheel useful. Clay is shaped to form a vessel;\n
What is not there makes the vessel useful. Doors and windows are cut to form a room;\n
What is not there makes the room useful.\n
Therefore, take advantage of what is there, By making use of what is not.\n
`
	passages[11] = `CONTROLLING THE SENSES\n
The five colors will blind one's eye. The five tones will deafen one's ear. The five flavors will jade one's taste.\n
Racing and hunting will derange one's mind. Goods that are hard to get will obstruct one's way.\n
Therefore, Evolved Individuals Regard the center and not the eye. Hence they discard one and receive the other.\n
`
	passages[12] = `EXPANDING IDENTIFICATION\n
There is alarm in both favor and disgrace. Esteem and fear are identified with the self.\n
What is the meaning of "alarm in both favor and disgrace?" Favor ascends; disgrace descends. To attain them brings alarm. To lose them brings alarm. That is the meaning of "alarm in both favor and disgrace."
What is the meaning of "esteem and fear are identified with the self?" The reason for our fear Is the presence of our self. When we are selfless, What is there to fear?\n
Therefore those who esteem the world as self\n
Will be committed to the world. Those who love the world as self\n
Will be entrusted with the world.\n
`
	passages[13] = `THE ESSENCE OF TAO\n
Looked at but not seen:\n
Its name is formless. Listened to but not heard:\n
Its name is soundless. Reached for but not obtained:\n
Its name is intangible.\n
These three cannot be analyzed, So they mingle and act as one.\n
Its rising is not bright;\n
Its setting is not dark. Endlessly, the nameless goes on,\n
Merging and returning to nothingness.\n
That is why it is called\n
The form of the formless,\n
The image of nothingness. That is why it is called elusive.\n
Confronted, its beginning is not seen. Followed, its end is not seen.\n
Hold on to the ancient Tao;\n
Control the current reality. Be aware of the ancient origins;\n
This is called the Essence of Tao.\n
`
	passages[14] = `THE POWER IN SUBTLE FORCE\n
Those skillful in the ancient Tao Are subtly ingenious and profoundly intuitive. They are so deep they cannot be recognized. Since, indeed, they cannot be recognized, Their force can be contained.\n
\n
So careful!\n
As if wading a stream in winter. So hesitant!\n
As if respecting all sides in the community. So reserved!\n
As if acting as a guest. So yielding!\n
As if ice about to melt. So candid!\n
As if acting with simplicity. So open!\n
As if acting as a valley. So integrated!\n
As if acting as muddy water.\n
Who can harmonize with muddy water,\n
And gradually arrive at clarity? Who can move with stability,\n
And gradually bring endurance to life?\n
Those who maintain the Tao\n
Do not desire to become full. Indeed, since they are not full,\n
They can be used up and also renewed.\n
`
	passages[15] = `KNOWING THE ABSOLUTE\n
Attain the highest openness;\n
Maintain the deepest harmony. Become a part of All Things;\n
In this way, I perceive the cycles.\n
Indeed, things are numerous;\n
But each cycle merges with the source. Merging with the source is called harmonizing;\n
This is known as the cycle of destiny.\n
The cycle of destiny is called the Absolute;\n
Knowing the Absolute is called insight. To not know the Absolute\n
Is to recklessly become a part of misfortune.\n
To know the Absolute is to be tolerant.\n
What is tolerant becomes impartial; What is impartial becomes powerful; What is powerful becomes natural; What is natural becomes Tao.\n
What has Tao becomes everlasting And free from harm throughout life.\n
`
	passages[16] = ``
	passages[17] = ``
	passages[18] = ``
	passages[19] = ``
	passages[20] = ``
	passages[21] = ``
	passages[22] = ``
	passages[23] = ``
	passages[24] = ``
	passages[25] = ``
	passages[26] = ``
	passages[27] = ``
	passages[28] = ``
	passages[29] = ``
	passages[30] = ``
	passages[31] = ``
	passages[32] = ``
	passages[33] = ``
	passages[34] = ``
	passages[35] = ``
	passages[36] = ``
	passages[37] = ``
	passages[38] = ``
	passages[39] = ``
	passages[40] = ``
	passages[41] = ``
	passages[42] = ``
	passages[43] = ``
	passages[44] = ``
	passages[45] = ``
	passages[46] = ``
	passages[47] = ``
	passages[48] = ``
	passages[49] = ``
	passages[50] = ``
	passages[51] = ``
	passages[52] = ``
	passages[53] = ``
	passages[54] = ``
	passages[55] = ``
	passages[56] = ``
	passages[57] = ``
	passages[58] = ``
	passages[59] = ``
	passages[60] = ``
	passages[61] = ``
	passages[62] = ``
	passages[63] = ``
	passages[64] = ``
	passages[65] = ``
	passages[66] = ``
	passages[67] = ``
	passages[68] = ``
	passages[69] = ``
	passages[70] = ``
	passages[71] = ``
	passages[72] = ``
	passages[73] = ``
	passages[74] = ``
	passages[75] = ``
	passages[76] = ``
	passages[77] = ``
	passages[78] = ``
	passages[79] = ``
	passages[80] = ``
	return passages[rand.Intn(80)]
}
