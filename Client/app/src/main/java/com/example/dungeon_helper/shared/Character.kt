package com.example.dungeon_helper.shared

import kotlinx.serialization.Serializable
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

@Serializable
data class Character(
    var hp: Int,
    var exp: Int,
    var avatar: String,
    var charName: String,
    var sex: Boolean,
    var weight: Int,
    var height: Int,
    var charClass: CharacterClass,
    var race: Race,
    var subrace: Subrace,
    var stats: Stats,
    var addLanguage: String,
    var characterSkills: List<CharacterSkill>,
    var alignment: Alignment,
    var ideals: String,
    var weaknesses: String,
    var traits: String,
    var allies: String,
    var organizations: String,
    var enemies: String,
    var story: String,
    var goals: String,
    var treasures: String,
    var notes: String
) {

    override fun toString(): String {
        return "Character(hp=$hp, exp=$exp, avatar='$avatar', charName='$charName', sex=$sex, weight=$weight, height=$height, charClass=$charClass, race=$race, subrace=$subrace, stats=$stats, addLanguage='$addLanguage', characterSkills=$characterSkills, alignment=$alignment, ideals='$ideals', weaknesses='$weaknesses', traits='$traits', allies='$allies', organizations='$organizations', enemies='$enemies', story='$story', goals='$goals', treasures='$treasures', notes='$notes')"
    }
}

@Serializable
data class GetCharacter(
    var id: Int,
    var hp: Int,
    var lvl: Int,
    var exp: Int,
    var avatar: String,
    var charName: String,
    var sex: Boolean,
    var weight: Int,
    var height: Int,
    var charClass: String,
    var race: String,
    var subrace: String,
    var stats: Stats,
    var addLanguage: String,
    var characterSkills: String,
    var alignment: String,
    var ideals: String,
    var weaknesses: String,
    var traits: String,
    var allies: String,
    var organizations: String,
    var enemies: String,
    var story: String,
    var goals: String,
    var treasures: String,
    var notes: String
) {

    override fun toString(): String {
        return "GetCharacter(id=$id, hp=$hp, lvl=$lvl, exp=$exp, avatar='$avatar', charName='$charName', sex=$sex, weight=$weight, height=$height, charClass='$charClass', race='$race', subrace='$subrace', stats=$stats, addLanguage='$addLanguage', characterSkills='$characterSkills', alignment='$alignment', ideals='$ideals', weaknesses='$weaknesses', traits='$traits', allies='$allies', organizations='$organizations', enemies='$enemies', story='$story', goals='$goals', treasures='$treasures', notes='$notes')"
    }
}

@Serializable
data class AllCharacter(
    var idChar: Int,
    var charName: String,
    var avatar: String
) {
    override fun toString(): String {
        return "GetCharacter(idChar=$idChar, charName='$charName', avatar='$avatar')"
    }
}

@Serializable
data class CharacterClass(
    var id: Int,
    var className: String
){
    override fun toString(): String {
        return "CharacterClass(id=$id, className='$className')"
    }
}

@Serializable
data class Race(
    var id: Int,
    var raceName: String
){
    override fun toString(): String {
        return "Race(id=$id, raceName='$raceName')"
    }
}

@Serializable
data class Subrace(
    var id: Int,
    var raceName: String,
    var stats: Stats
){
    override fun toString(): String {
        return "Subrace(id=$id, raceName='$raceName', stats=$stats)"
    }
}

@Serializable
data class Stats(
    var strength: Int,
    var dexterity: Int,
    var constitution: Int,
    var intelligence: Int,
    var wisdom: Int,
    var charisma: Int
) {
    override fun toString(): String {
        return "Stats(strength=$strength, dexterity=$dexterity, constitution=$constitution, intelligence=$intelligence, wisdom=$wisdom, charisma=$charisma)"
    }
}

@Serializable
data class CharacterSkill(
    var id: Int,
    var skillName: String
){
    override fun toString(): String {
        return "CharacterSkill(id=$id, skillName='$skillName')"
    }
}

@Serializable
data class Alignment(
    var id: Int,
    var alignmentName: String
){
    override fun toString(): String {
        return "Alignment(id=$id, alignmentName='$alignmentName')"
    }
}

object JsonHelper {
    private val json = Json { prettyPrint = true }


    fun serializeToJson(character: Character): String {
        return json.encodeToString(character)
    }

    fun deserializeFromJsonCharacter(jsonString: String): Character {
        return json.decodeFromString(jsonString)
    }

    fun serializeToJson(character: GetCharacter): String {
        return json.encodeToString(character)
    }

    fun deserializeFromJsonGetCharacter(jsonString: String): GetCharacter {
        return json.decodeFromString(jsonString)
    }


    fun serializeToJson(characterClass: CharacterClass): String {
        return json.encodeToString(characterClass)
    }

    fun deserializeFromJsonClass(jsonString: String): CharacterClass {
        return json.decodeFromString(jsonString)
    }


    fun serializeToJson(race: Race): String {
        return json.encodeToString(race)
    }

    fun deserializeFromJsonRace(jsonString: String): Race {
        return json.decodeFromString(jsonString)
    }


    fun serializeToJson(subrace: Subrace): String {
        return json.encodeToString(subrace)
    }

    fun deserializeFromJsonSubrace(jsonString: String): Subrace {
        return json.decodeFromString(jsonString)
    }


    fun serializeToJson(stats: Stats): String {
        return json.encodeToString(stats)
    }

    fun deserializeFromJsonStats(jsonString: String): Stats {
        return json.decodeFromString(jsonString)
    }


    fun serializeToJson(characterSkill: CharacterSkill): String {
        return json.encodeToString(characterSkill)
    }

    fun deserializeFromJsonCharacterSkill(jsonString: String): CharacterSkill {
        return json.decodeFromString(jsonString)
    }


    fun serializeToJson(alignment: Alignment): String {
        return json.encodeToString(alignment)
    }

    fun deserializeFromJsonAlignment(jsonString: String): Alignment {
        return json.decodeFromString(jsonString)
    }



//    fun serializeListToJson(characterClasses: List<CharacterClass>): String {
//        return json.encodeToString(characterClasses)
//    }

    fun deserializeListFromJsonCharacterClassList(jsonString: String): List<CharacterClass> {
        return json.decodeFromString(jsonString)
    }


//    fun serializeListToJson(races: List<Race>): String {
//        return json.encodeToString(races)
//    }

    fun deserializeListFromJsonRaceList(jsonString: String): List<Race> {
        return json.decodeFromString(jsonString)
    }


//    fun serializeListToJson(subraces: List<Subrace>): String {
//        return json.encodeToString(subraces)
//    }

    fun deserializeListFromJsonSubraceList(jsonString: String): List<Subrace> {
        return json.decodeFromString(jsonString)
    }


    fun deserializeListFromJsonAlignmentList(jsonString: String): List<Alignment> {
        return json.decodeFromString(jsonString)
    }

    fun deserializeListFromJsonSkillList(jsonString: String): List<CharacterSkill> {
        return json.decodeFromString(jsonString)
    }

    fun deserializeListFromJsonAllCharacter(jsonString: String): List<AllCharacter> {
        return json.decodeFromString(jsonString)
    }


}

