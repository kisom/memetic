package poem

var rubaiyat = Poem{
	Title:  "The Rubaiyat of Omar Khayyam",
	Author: "Edward Fitzgerald",
	Year:   1889,
	Stanzas: readStanzas(`WAKE! For the Sun, who scatter'd into flight
 The Stars before him from the Field of Night,
   Drives Night along with them from Heav'n, and strikes
 The Sultan's Turret with a Shaft of Light.

 Before the phantom of False morning died,
 Methought a Voice within the Tavern cried,
   "When all the Temple is prepared within,
 "Why nods the drowsy Worshiper outside?"

 And, as the Cock crew, those who stood before
 The Tavern shouted--"Open then the Door!
   "You know how little while we have to stay,
 And, once departed, may return no more."

 Now the New Year reviving old Desires,
 The thoughtful Soul to Solitude retires,
   Where the WHITE HAND OF MOSES on the Bough
 Puts out, and Jesus from the Ground suspires.

 Iram indeed is gone with all his Rose,
 And Jamshyd's Sev'n-ring'd Cup where no one knows;
   But still a Ruby kindles in the Vine,
 And many a Garden by the Water blows.

 And David's lips are lockt; but in divine
 High-piping Pehlevi, with "Wine! Wine! Wine!
   "Red Wine!"--the Nightingale cries to the Rose
 That sallow cheek of hers to' incarnadine.

 Come, fill the Cup, and in the fire of Spring
 Your Winter garment of Repentance fling:
   The Bird of Time has but a little way
 To flutter--and the Bird is on the Wing.

 Whether at Naishapur or Babylon,
 Whether the Cup with sweet or bitter run,
   The Wine of Life keeps oozing drop by drop,
 The Leaves of Life keep falling one by one.

 Each Morn a thousand Roses brings, you say:
 Yes, but where leaves the Rose of Yesterday?
   And this first Summer month that brings the Rose
 Shall take Jamshyd and Kaikobad away.

 Well, let it take them!  What have we to do
 With Kaikobad the Great, or Kaikhosru?
   Let Zal and Rustum bluster as they will,
 Or Hatim call to Supper--heed not you.

 With me along the strip of Herbage strown
 That just divides the desert from the sown,
   Where name of Slave and Sultan is forgot--
 And Peace to Mahmud on his golden Throne!

 A Book of Verses underneath the Bough,
 A Jug of Wine, a Loaf of Bread--and Thou
   Beside me singing in the Wilderness--
 Oh, Wilderness were Paradise enow!

 Some for the Glories of This World; and some
 Sigh for the Prophet's Paradise to come;
   Ah, take the Cash, and let the Credit go,
 Nor heed the rumble of a distant Drum!

 Look to the blowing Rose about us--"Lo,
 Laughing," she says, "into the world I blow,
   At once the silken tassel of my Purse
 Tear, and its Treasure on the Garden throw."

 And those who husbanded the Golden grain,
 And those who flung it to the winds like Rain,
   Alike to no such aureate Earth are turn'd
 As, buried once, Men want dug up again.

 The Worldly Hope men set their Hearts upon
 Turns Ashes--or it prospers; and anon,
   Like Snow upon the Desert's dusty Face,
 Lighting a little hour or two--is gone.

 Think, in this batter'd Caravanserai
 Whose Portals are alternate Night and Day,
   How Sultan after Sultan with his Pomp
 Abode his destined Hour, and went his way.

 They say the Lion and the Lizard keep
 The courts where Jamshyd gloried and drank deep:
   And Bahram, that great Hunter--the Wild Ass
 Stamps o'er his Head, but cannot break his Sleep.

 I sometimes think that never blows so red
 The Rose as where some buried Caesar bled;
   That every Hyacinth the Garden wears
 Dropt in her Lap from some once lovely Head.


 And this reviving Herb whose tender Green
 Fledges the River-Lip on which we lean--
   Ah, lean upon it lightly! for who knows
 From what once lovely Lip it springs unseen!
`),
}

func init() {
	register(rubaiyat)
}
